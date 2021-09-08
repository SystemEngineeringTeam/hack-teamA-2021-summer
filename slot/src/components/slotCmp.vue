<template>
  <section class="panel">
    <img v-bind:src="image" v-bind:class="{ inactive: isUnmatched }" />
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
      ],
      image: require("@/assets/7.jpg"),
      timeoutId: "",
      isSelected: true,
      isUnmatched: false,
    };
  },
  // 枠組みの中のHTMl制作
  template: `<section class="panel">
                        <img v-bind:src= image v-bind:class={inactive:isUnmatched}>
                        <div class="stop" v-on:click="stop()" 
                        v-bind:class={inactive:isSelected}>STOP</div>
                    </section>`,
  methods: {
    // ランダムな画像を表示
    getRandomImage() {
      this.image = this.images[Math.floor(Math.random() * this.images.length)];
    },
    getImage() {
        if(this.image == this.images[0]){
            this.image = this.images[1];
        }
        else if(this.image == this.images[1]){
            this.image = this.images[2];
        }
        else if(this.image == this.images[2]){
            this.image = this.images[3];
        }
        else if(this.image == this.images[3]){
            this.image = this.images[0];
        }
        // for(let i = 0; i < this.images.length; i ++){
        //     if(i < this.images.length - 1){
        //         if(this.image == this.images[i]){
        //             this.image = this.images[i+1];
        //         }
        //     }
        //     else {
        //         if(i == this.images.length){
        //             this.image = this.images[0];
        //         }
        //     }
        // }
    },
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

  width: 80px;
  height: 80px;
  margin-left: 55px;
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
  width: 200px;
  height: 200px;
  margin-top: 20px;
  margin-bottom: 70px;
}
</style>
