<template>
  <div id="app">
    <div class="flexbox flexbox-left">
      <slotComponent v-on:sendImg="receiveImg" />
      <img v-for="image in this.childImg" :src="image" v-bind:key="image.id" />
    </div>
    <main>
      <div class="dis">
        <!-- v-on押したら実行される  -->
        <slot-component
          ref="component1"
          v-on:decrement="decrementPanel"
        ></slot-component>
        <slot-component
          ref="component2"
          v-on:decrement="decrementPanel"
        ></slot-component>
        <slot-component
          ref="component3"
          v-on:decrement="decrementPanel"
        ></slot-component>
      </div>

      <div class="sub">
        コイン枚数
        {{ this.coment }}
        {{ this.coin }}
        <a
          href="#"
          class="btn-sf-like"
          v-on:click="reset()"
          v-bind:class="{ inactive: isRunning }"
        >
          コイン リセット
        </a>
        <!-- <v-btn @click="$router.push('settting')">画像選択するよ</v-btn> -->
        <div
          class="spin"
          v-on:click="spin()"
          v-bind:class="{ inactive: isRunning }"
        >
          MAX
        </div>
      </div>
    </main>
    <v-btn>
      <router-link to="/setting" style="text-decoration:" none
        >絵柄変更</router-link
      >
    </v-btn>
    <chart v-if="showChart"></chart>
  </div>
</template>
<style scoped>
body {
  background: #bdc3c7;
  font-size: 16px;
  font-weight: bold;
  font-family: Arial, Helvetica, sans-serif;
}
.sub {
  font-size: 30px;
  color: #ffffff;
  width: 200px;
  height: 300px;
  padding: 20px;
  /* border-radius: 12px; */
}

.dis {
  width: 800px;
  height: 350px;
  background: #dab300;
  padding: 20px;
  border-radius: 12px;
  margin: 16px auto;
  display: flex;
  justify-content: space-between;
}
.btn-sf-like {
  position: relative;
  display: inline-block;
  font-weight: bold;
  text-decoration: none;
  color: #fff;
  text-shadow: 0 0 5px rgba(255, 255, 255, 0.73);
  padding: 0.3em 0.5em;
  background: #00bcd4;
  border-top: solid 3px #00a3d4;
  border-bottom: solid 3px #00a3d4;
  transition: 0.4s;
  font-size: 15px;
  margin-top: 50px;
}

.btn-sf-like:hover {
  text-shadow: -6px 0px 15px rgba(255, 255, 240, 0.83),
    6px 0px 15px rgba(255, 255, 240, 0.83);
}
.flexbox {
  display: flex;
  flex-direction: column;
}

.btn-sf-like2 {
  width: 400px;
  position: relative;
  display: inline-block;
  font-weight: bold;
  text-decoration: none;
  color: #fff;
  text-shadow: 0 0 5px rgba(255, 255, 255, 0.73);
  padding: 0.3em 0.5em;
  background: #00bcd4;
  border-top: solid 3px #00a3d4;
  border-bottom: solid 3px #00a3d4;
  transition: 0.4s;
  font-size: 40px;
  /* margin-top: 50px; */
}

.btn-sf-like2:hover {
  text-shadow: -6px 0px 15px rgba(255, 255, 240, 0.83),
    6px 0px 15px rgba(255, 255, 240, 0.83);
}

main {
  width: 1100px;
  height: 500px;
  background: rgb(178, 34, 34);
  padding: 20px;
  /* border: 50px solid #fff; */
  border-radius: 12px;
  margin: 16px auto;
  display: flex;
  justify-content: space-between;
}

input {
  display: none; /* デフォルトのボタンを非表示にする */
}

.spin {
  width: 100px;
  height: 50px;
  background: #008cff;
  box-shadow: 0 4px 0 #2880b9;
  border-radius: 18px;
  line-height: 50px;
  text-align: center;
  font-size: 20px;
  color: rgb(255, 255, 255);
  user-select: none;
  margin-top: 100px;
  color: #fff; /* 文字色を白に指定 */
  background: rgb(2, 133, 255); /* 背景色をオレンジに指定 */
  animation: flash 1s infinite; /* アニメーションflashを1秒ごとに繰り返す */
}

@keyframes flash {
  0%,
  100% {
    /* 明るく光るよう影を重ねる */
    box-shadow: 0 0 10px #ffc, 0 0 20px #ffc, 0 0 30px #ff9, 0 0 40px #ff6,
      0 0 70px #fc6, 0 0 80px #f99, 0 0 100px #ff96, 0 0 150px #ff96;
  }
  50% {
    /* 淡く光るよう影を重ねる */
    box-shadow: 0 0 10px #fff, 0 0 20px #fff, 0 0 30px #ffc, 0 0 40px #ff9;
  }
}
</style>
<script>
import slotComponent from "@/components/slotCmp.vue";
import Chart from "../components/Chart.vue";
// import App from "@/views/App.vue";
export default {
  // el: '#app',
  data() {
    return {
      panelLeft: 3,
      isRunning: false,
      isNotAllCorrect: true,
      isR: true,
      coin: 0,
      coment: "",
      spinCnt: 0,
      showChart: true,
      childImg: [],
    };
  },
  // 'slot-component'：javaでいうimportと同じ
  components: {
    "slot-component": slotComponent,
    Chart: Chart,
  },

  async created() {
    try {
      var res = await this.axios.get("http://localhost:8080/user", {
        headers: {
          Authorization: "Bearer " + this.$cookie.get("token"),
        },
      });
      var coin = res.data.coin;
      var spin = res.data.spin;
      this.coin = coin;
      this.spinCnt = spin;
    } catch (e) {
      console.log(e);
    }
  },
  methods: {
    async sendChart() {
      this.spinCnt += 1;
      try {
        var res = await this.axios.post(
          "http://localhost:8080/chartdata",
          {
            spin: this.spinCnt,
            coin: this.coin,
          },
          {
            headers: {
              Authorization: "Bearer " + this.$cookie.get("token"),
            },
          }
        );
        console.log(res);
        this.update();
      } catch (e) {
        console.log(e);
      }
    },
    receiveImg(images) {
      this.childImg = images;
      console.log("hoge", images);
    },
    total() {
      this.axios
        //   今回はポストに接続するので.postにする。第一引数に宛先を指定、第二引数に送りたいデータを指定する
        .post(
          "http://localhost:8080/coin",
          {
            total: this.coin,
          },
          {
            headers: {
              Authorization: "Bearer " + this.$cookie.get("token"),
            },
          }
        )
        // 成功した場合（.then)console.logにresponseを返す
        .then((response) => {
          console.log(response);
        })
        // 失敗した場合(.catch)console.logにerrorを返す
        .catch((error) => {
          console.log(error);
        });
    },
    // 押したときに実行、
    spin() {
      if (this.isRunning) {
        return;
      }
      //isRunningfalseの時isRunningをtrueにして回転させる
      this.isRunning = true;
      if (this.isR) {
        this.coin -= 3;
      }
      this.isR = true;

      this.$refs.component1.activate();
      this.$refs.component2.activate();
      this.$refs.component3.activate();

      this.$refs.component1.spin();
      this.$refs.component2.spin();
      this.$refs.component3.spin();
    },
    reset() {
      if (this.isRunning) {
        return;
      }
      this.coin = 100;
    },
    update() {
      this.showChart = false;
      this.$nextTick(() => {
        this.showChart = true;
      });
    },
    // 実行されたらpanelLeftが１ずつマイナスされていく。panelLeftが０になるとthis.isRunning = falseにし、回転を止める
    decrementPanel: function() {
      this.panelLeft--;
      if (this.panelLeft == 0) {
        this.isRunning = false;
        this.panelLeft = 3;
        if (
          this.$refs.component1.images[this.$refs.component1.image] !==
            this.$refs.component2.images[this.$refs.component2.image] &&
          this.$refs.component1.images[this.$refs.component1.image] !==
            this.$refs.component3.images[this.$refs.component3.image]
        ) {
          this.$refs.component1.isUnmatched = true;
        }
        if (
          this.$refs.component2.images[this.$refs.component2.image] !==
            this.$refs.component1.images[this.$refs.component1.image] &&
          this.$refs.component2.images[this.$refs.component2.image] !==
            this.$refs.component3.images[this.$refs.component3.image]
        ) {
          this.$refs.component2.isUnmatched = true;
        }
        if (
          this.$refs.component3.images[this.$refs.component3.image] !==
            this.$refs.component1.images[this.$refs.component1.image] &&
          this.$refs.component3.images[this.$refs.component3.image] !==
            this.$refs.component2.images[this.$refs.component2.image]
        ) {
          this.$refs.component3.isUnmatched = true;
        }
        if (
          this.$refs.component1.images[this.$refs.component1.image] ==
            this.$refs.component2.images[this.$refs.component2.image] &&
          this.$refs.component1.images[this.$refs.component1.image] ==
            this.$refs.component3.images[this.$refs.component3.image] &&
          this.$refs.component2.images[this.$refs.component2.image] ==
            this.$refs.component3.images[this.$refs.component3.image]
        ) {
          if (this.$refs.component1.image == 0) {
            this.coin += 300;
          } else if (this.$refs.component1.image == 1) {
            this.coin += 200;
          } else if (
            this.$refs.component1.image == 2 ||
            this.$refs.component1.image == 4
          ) {
            this.coin += 100;
          } else if (
            this.$refs.component1.image == 3 ||
            this.$refs.component1.image == 5
          ) {
            this.isR = false;
          }
        }
        this.total();
        this.sendChart();
      }
    },
  },
};
</script>
