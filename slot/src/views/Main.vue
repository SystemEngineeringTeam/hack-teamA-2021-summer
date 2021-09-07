<template>
    <div id="app">
        <body>
        <main>
            <slot-component ref="component1" v-on:decrement="decrementPanel"></slot-component>
            <slot-component ref="component2" v-on:decrement="decrementPanel"></slot-component>
            <slot-component ref="component3" v-on:decrement="decrementPanel"></slot-component>
        </main>
        </body>
        <div></div>
        <div class="spin" v-on:click="spin()" v-bind:class={inactive:isRunning}>SPIN</div> 
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

        main {
            width: 900px;
            height: 400px;
            background: #000000;
            padding: 20px;
            /* border: 50px solid #fff; */
            border-radius: 12px;
            ;
            margin: 16px auto;
            display: flex;
            justify-content: space-between;
        }

        .panel img {
            width: 200px;
            height: 150px;
            margin-top: 20px;
            margin-bottom: 100px;
        }

        .stop {
            cursor: pointer;
            width: 90px;
            height: 32px;
            background: #ef454a;
            box-shadow: 0 4px 0 #d1483e;
            border-radius: 16px;
            line-height: 32px;
            text-align: center;
            font-size: 14px;
            color: #fff;
            user-select: none;
        }

        .spin {
            cursor: pointer;
            width: 280px;
            height: 36px;
            background: #3498db;
            box-shadow: 0 4px 0 #2880b9;
            border-radius: 18px;
            line-height: 36px;
            text-align: center;
            font-size: 14px;
            color: #fff;
            user-select: none;
            margin: 0 auto;
        }

        .unmatched {
            opacity: 0.5;
        }

        .inactive {
            opacity: 0.5;
        }
</style>
<script>
// import slotComponent from '@/components/slotCmp.vue'
    export default {
            el: '#app',
            data() {
                return {
                    panelLeft: 3,
                    isRunning: false,
                    isNotAllCorrect: true,
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
                    }
                },

            }
        };
</script>
