<template>
    <div id="app">
      <main>
        <div class = 'dis'>
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
          
        <sub>
           コイン枚数
           {{this.coin}}
           <div
          class="spin"
          v-on:click="spin()"
          v-bind:class="{inactive:isRunning}"
        >
          MAX
        </div>
        </sub>
           
      </main>
    </div>
</template>
    <script src="https://cdn.jsdelivr.net/npm/vue/dist/vue.js"></script> 
<style scoped>
        body {
            background: #bdc3c7;
            font-size: 16px;
            font-weight: bold;
            font-family: Arial, Helvetica, sans-serif;
        }
        sub{
            fontsize: 20px;
            color: #ffffff;
            width: 100px;
            height: 50px;
            padding: 20px;
            border-radius: 12px;
            
        }
        .dis{
            width: 700px;
            height: 250px;
            background: #dab300;
            padding: 20px;
            border-radius: 12px;
            ;
            margin: 16px auto;
            display: flex;
            justify-content: space-between;
        }
        main {
            width: 900px;
            height: 400px;
            background: rgb(178,34,34);
            padding: 20px;
            /* border: 50px solid #fff; */
            border-radius: 12px;
            ;
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
    margin-top: 200px;
  color: #fff;        /* 文字色を白に指定 */ 
  background: rgb(2, 133, 255);   /* 背景色をオレンジに指定 */
  animation: flash 1s infinite; /* アニメーションflashを1秒ごとに繰り返す */
}

@keyframes flash {
  0%, 100% {
    /* 明るく光るよう影を重ねる */
    box-shadow: 0 0 10px #ffc, 0 0 20px #ffc, 0 0 30px #ff9, 0 0 40px #ff6, 0 0 70px #fc6, 0 0 80px #f99, 0 0 100px #ff96, 0 0 150px #ff96;
  }
  50% {
    /* 淡く光るよう影を重ねる */
    box-shadow: 0 0 10px #fff, 0 0 20px #fff, 0 0 30px #ffc, 0 0 40px #ff9;
  }
}

        .unmatched {
            opacity: 0.5;
        }

        .inactive {
            opacity: 0.5;
        }

    </style>
<script>
import slotComponent from '@/components/slotCmp.vue'
    export default {
            // el: '#app',
            data() {
                return {
                    panelLeft: 3,
                    isRunning: false,
                    isNotAllCorrect: true,
                    coin: 100,
                }
            },
            // 'slot-component'：javaでいうimportと同じ
            components: {
                'slot-component': slotComponent
            },
            methods: {
                spin() {
                    if (this.isRunning) {
                        return;
                    }
                    //isRunningfalsenの時isRunningをtrueにして
                    this.isRunning = true;
                    this.coin -= 3;
                    this.$refs.component1.activate();
                    this.$refs.component2.activate();
                    this.$refs.component3.activate();

                    this.$refs.component1.spin();
                    this.$refs.component2.spin();
                    this.$refs.component3.spin();
                },

                decrementPanel: function () {
                    this.panelLeft--;
                    if (this.panelLeft == 0) {
                        this.isRunning = false;
                        this.panelLeft = 3;

                        if (this.$refs.component1.image !== this.$refs.component2.image && 
                            this.$refs.component1.image !== this.$refs.component3.image) {

                            this.$refs.component1.isUnmatched = true;
                        }
                        if (this.$refs.component2.image !== this.$refs.component1.image && 
                            this.$refs.component2.image !== this.$refs.component3.image) {

                            this.$refs.component2.isUnmatched = true;

                        }
                        if (this.$refs.component3.image !== this.$refs.component1.image && 
                            this.$refs.component3.image !== this.$refs.component2.image) {

                            this.$refs.component3.isUnmatched = true;
                        }
                        if (this.$refs.component1.image == this.$refs.component2.image && 
                            this.$refs.component1.image == this.$refs.component3.image) {
                        }

                        // if(this.$refs.component1.image == this.$refs.component2.image == this.$refs.component3.image){
                        //     for(let i = 0; i < this.data.image.length; i++){
                        //         if(i == 1){
                        //             this.coin += 300;
                        //         }
                        //         else if(i == 2){
                        //             this.coin += 150;
                        //         }
                        //     }
                        // }
                    }
                },
            }
        };
</script>
