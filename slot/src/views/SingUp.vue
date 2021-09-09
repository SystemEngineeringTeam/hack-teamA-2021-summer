<template>
  <div>
    <body>
        <div class="header">
        <div class="header-logo">ようこそ賭博へ!</div>
        </div>

        <div class="main">
        <div class="copy-container">
            <h1>闇落ちの儀式</h1>
        </div>
        
        <div class="contents">
            
            <div class="contents-item">
                <p class="name">お名前</p>
                <input class="names" v-model="name"/>
                <p class="mail">メールアドレス</p>
                <input class="mails" v-model="mail"/>
                <p class="pass">パスワード</p>
                <input class="pases" v-model="password"/>
                
            </div>
            <img class="bill" src="https://tsukatte.com/wp-content/uploads/2020/08/money-with-wings.png">
            <img class="bills" src="https://tsukatte.com/wp-content/uploads/2020/08/money-with-wings.png">
        </div>

        <button @click="singup()">登録</button>
        
        <div class="contact-form">
        </div>
        </div>
    </body>
    </div>    
</template>

<script>
export default {
  data: () => {
    return {
      name:"",
      mail: "",
      password: "",
    };
  },
  methods: {
    singup() {
      this.axios
        .post("http://localhost:8080/user", {
          name: this.name,
          email: this.mail,
          password: this.password,
        },)
        .then((response) => {
          console.log(response);
        })
        .catch((error) => {
          console.log(error);
        });
    },
  },
    actions: {},
};
  
</script>

<style scoped>
body {
    position: relative;
    padding:0;
    font-family: "Avenir Next";
    background-color: #000000;
    background-image:url("https://www.pakutaso.com/shared/img/thumb/PAK86_hikaritobu420150318210141_TP_V.jpg")
    
}

li {
    list-style: none;
}

/*headerの設定、背景の色*/
.header {
    position: relative;
    padding:0;
    background-color: #b22222;    color: #33FF33;
    height: 300px;
    background-image:url("https://chicodeza.com/wordpress/wp-content/uploads/okane-illust6.png")
}

/*heder、チーム名とかロゴ的な何かの大きさ、位置*/
.header-logo {
    float: left;
    font-size: 120px;
    padding: 20px 40px;
    position:absolute;
    top: 0;
    left: 0;
}

/*headerおわり*/
/*mainの設定*/
.main {
    padding: 100px 80px;
}

.copy-container h1 {
    font-size: 100px;
    color: #DD0000;
    position: absolute;
    top: 220px;
    left: 10px;
}

.name {
    color: #fff;
    font-size: 30px;
    position:absolute;
    top: 500px;
    left: 80px;
}
.mail {
    color: #fff;
    position:absolute;
    top: 600px;
    left: 80px;
}
.pass {
    color: #fff;
    position:absolute;
    top: 700px;
    left: 80px;
}

.names{
    border: 4px solid #000;
    width: 400px;
    height: 50px;
    font-size: 30px;
    background-color: #4b0082;
    color: #fff;
    position:absolute;
    top: 560px;
    left: 80px;
}

.mails{
    border: 4px solid #000;
    width: 400px;
    height: 50px;
    font-size: 30px;
    background-color: #4b0082;
    color: #fff;
    position:absolute;
    top: 660px;
    left: 80px;
}

.pases{
    border: 4px solid #000;
    width: 400px;
    height: 50px;
    font-size: 30px;
    background-color: #4b0082;
    color: #fff;
    position:absolute;
    top: 760px;
    left: 80px;
}

.bill{
    width: 300px;
    height: 300px;
    position:absolute;
    bottom: 100px;
    right: 50px;
}
.bills{
    width: 300px;
    height: 300px;
    position:absolute;
    bottom: 400px;
    right: 180px;
}

button{
    width: 120px;
    height: 80px;
    font-size: 50px;
    color: #b22222;
    background-color: #4b0082;
    position:absolute;
    bottom: 150px;
    left: 220px;

}

.contents {
    height: 500px;
    
    margin-top: 100px;
    
}

.section-title {
    border-bottom: 2px solid #dee7ec;
    
    font-size: 28px;
    
    padding-bottom: 15px;
    
    margin-bottom: 50px;
}

.contents-item {
    float: left;
    
    margin-right: 40px;
    
}

.contents-item p{
    font-size: 24px;
    margin-top: 30px;
}


</style>