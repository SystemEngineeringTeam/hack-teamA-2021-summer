<script>
        var slotComponent = Vue.extend({
            data() {
                return {
                    //スロットの絵柄
                    images: [
            'slot_img/7.jpg',
            'slot_img/b.jpg',
            'slot_img/7.jpg',
        ],
        image: 'slot_img/7.jpg',
                    timeoutId: '',
                    isSelected: true,
                    isUnmatched: false,
                }
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
                // 回る時の処理
                spin() {
                    // settimeout: 第二引数に与えられた実行タイミング(ミリ秒)で、第一引数に定義された処理内容を1度実行する。
                    // ここでは第一引数でthis.spin();を定義しているので永遠に並び続ける
                    this.timeoutId = setTimeout(() => {
                        this.getRandomImage();
                        this.spin();
                    }, 10);
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
                    this.$emit('decrement');
                },
                activate() {
                    this.isSelected = false;
                    this.isUnmatched = false;
                }
            },
        });