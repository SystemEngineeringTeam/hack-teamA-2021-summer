<template>
  <section class="panel">
    <!-- trueだったら半透明にする -->
    <img v-bind:src="images[image]" v-bind:class="{ inactive: isUnmatched }" />
    <div
      class="stop"
      v-on:click="stop()"
      v-bind:class="{ inactive: isSelected }"
    >
      STOP
    </div>
  </section>
</template>
<script>
export default {
  name: "slotCmp",
  data() {
    return {
      //スロットの絵柄
      images: [
        require("@/assets/7.jpg"),
        require("@/assets/B.jpg"),
        require("@/assets/c.jpg"),
        require("@/assets/r.jpg"),
        require("@/assets/c.jpg"),
        require("@/assets/r.jpg"),
      ],
      //   image: Math.floor(Math.random() * this.images.length),
      image: 0,
      timeoutId: "",
      isSelected: true,
      isUnmatched: false,
    };
  },
  async created() {
    try {
      var res = await this.axios.get("http://localhost:8080/setting", {
        headers: {
          Authorization: "Bearer " + this.$cookie.get("token"),
        },
      });
      console.log(res);
      var images = res.data.setting.split(":");
      for (let i = 0; i < images.length; i++) {
        images[i] =
          "http://localhost:8080/static/img/set" + (i + 1) + "/" + images[i];
      }
      images.push(images[2], images[3]);
      this.images = images;
    } catch (e) {
      console.log(e);
    }
    this.$emit("sendImg", this.images);
  },
  methods: {
    // ランダムな画像を表示
    // getRandomImage() {
    //   this.image = this.images[Math.floor(Math.random() * this.images.length)];
    // },
    getImage() {
      this.image = (this.image + 1) % this.images.length;
      // this.image = 3;
    },

    //     for(let i = 0; i < this.imageslength; i ++){
    //         if(i < this.images.length - 1){
    //             if(this.image == this.images[i]){
    //                 this.image = this.images[i+1];
    //             }
    //         }
    //         else {
    //             if(i == this.images.length){
    //                 this.image = this.images[0];
    //             }
    //         }
    //     }
    // },
    // 回る時の処理
    spin() {
      // settimeout: 第二引数に与えられた実行タイミング(ミリ秒)で、第一引数に定義された処理内容を1度実行する。
      // ここでは第一引数でthis.spin();を定義しているので永遠に並び続ける
      this.timeoutId = setTimeout(() => {
        // this.getRandomImage();
        this.getImage();
        this.spin();
      }, 50);
    },
    // ストップ時のメゾット
    stop() {
      if (this.isSelected) {
        return;
      }
      // isSelectedがtrueじゃない時
      // isSelectedをtrueにして
      // spinを止める
      this.isSelected = true;
      clearTimeout(this.timeoutId);
      // emit:子コンポーネントから親コンポーネントに値を渡す
      this.$emit("decrement");
    },
    activate() {
      this.isSelected = false;
      this.isUnmatched = false;
    },
  },
};
</script>
<style>
.stop {
  text-decoration: none;
  color: #ffffff;
  /* margin-top: 20px; */
  width: 80px;
  height: 80px;
  margin-left: 75px;
  line-height: 34px;
  font-size: 40px;
  border-radius: 100%;
  text-align: center;
  overflow: hidden;
  font-weight: bold;
  background-image: linear-gradient(#ffffff 0%, #ffffff 100%);
  text-shadow: 1px 1px 1px #dab300(255, 0, 0, 0.66);
  box-shadow: inset 0 2px 0 #dab300(255, 0, 0, 0.5),
    0 2px 2px rgba(0, 0, 0, 0.19);
  border-bottom: solid 2px #dab300;
  user-select: none;
}
.panel img {
  width: 240px;
  height: 250px;
  margin-top: 20px;
  margin-bottom: 70px;
}

.unmatched {
  opacity: 0.5;
}
.inactive {
  opacity: 0.5;
}
</style>
