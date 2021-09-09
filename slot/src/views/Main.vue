<template>
    <div id="app">
      <main>
        <div class = 'dis'>
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
          
        <div class = 'sub'>
           コイン枚数
           {{this.coment}}
           {{this.coin}}
           <a href="#" class="btn-sf-like" 
           v-on:click="reset()"
           v-bind:class="{inactive:isRunning}">
           コイン
           リセット
            </a>
           <div
          class="spin"
          v-on:click="spin()"
          v-bind:class="{inactive:isRunning}"
          >
            MAX
          </div>
        </div>
           
      </main>
      <a href="#" class="btn-sf-like"
      @click = total()>セーブ</a>
    </div>
</template>
<style scoped>
        body {
            background: #bdc3c7;
            font-size: 16px;
            font-weight: bold;
            font-family: Arial, Helvetica, sans-serif;
        }
        .sub{
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
        .btn-sf-like {
            position: relative;
            display: inline-block;
            font-weight: bold;
            text-decoration: none;
            color: #FFF;
            text-shadow: 0 0 5px rgba(255, 255, 255, 0.73);
            padding: 0.3em 0.5em;
            background: #00bcd4;
            border-top: solid 3px #00a3d4;
            border-bottom: solid 3px #00a3d4;
            transition: .4s;
            font-size: 20px;
}

        .btn-sf-like:hover {
        text-shadow: -6px 0px 15px rgba(255, 255, 240, 0.83),
                    6px 0px 15px rgba(255, 255, 240, 0.83);
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
                    isR:true,
                    coin: 100,
                    coment: "",
                }
            },
            // 'slot-component'：javaでいうimportと同じ
            components: {
                'slot-component': slotComponent
            },
            methods: {
                total() {
            this.axios
            //   今回はポストに接続するので.postにする。第一引数に宛先を指定、第二引数に送りたいデータを指定する
                .post("http://localhost:8080/coin", {
                total: this.coin,
                password: this.password,
                },)
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
                    if(this.isR){
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
                reset(){
                    if(this.isRunning){
                        return;
                    }
                    this.coin = 100;
                },
                // 実行されたらpanelLeftが１ずつマイナスされていく。panelLeftが０になるとthis.isRunning = falseにし、回転を止める
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
                            this.$refs.component1.image == this.$refs.component3.image && 
                            this.$refs.component2.image == this.$refs.component3.image) {
                                switch(this.$refs.component1.image){
                                    case 0:
                                        // <p>+300</p>
                                        // this.coment = "+300";
                                        this.coin += 300;
                                        break;
                                    
                                    case 1:
                                        // this.coment = "+200";
                                        this.coin += 200;
                                        break;

                                    case 2:
                                        // <p>+100</p>
                                        // this.coment = "+100";
                                        this.coin += 100;
                                        break;

                                    case 3:
                                        // this.coment = "もう１度";
                                        this.isR = false;
                                        break;
                                        
                                }

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
