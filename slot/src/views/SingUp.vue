<template>
  <div>
    <body>
        <div class="header">
        <div class="header-logo">卍亜琉派華潘堕合法賭博遊戯卍</div>
        </div>

        <div class="main">
        <div class="copy-container">
            <h1>ようこそ賭博へ</h1>
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
    position:relative;
    width: 100%;
    height: 100%;
    font-family: "Avenir Next";
    background-color: #000000;
    background-image:url("https://www.pakutaso.com/shared/img/thumb/PAK86_hikaritobu420150318210141_TP_V.jpg")
}

li {
    list-style: none;
}

/*headerの設定、背景の色*/
.header {
    position:relative;
    background-color: #b22222;    
    color: #33FF33;
    height: 240px;
    background-image:url("https://chicodeza.com/wordpress/wp-content/uploads/okane-illust6.png")
}

/*heder、チーム名とかロゴ的な何かの大きさ、位置*/
.header-logo {
    float: left;
    font-size: 60px;
    padding: 20px 40px;
    position: absolute;
    top: 20px;
    left: 30px;
}

/*headerおわり*/
/*mainの設定*/
.main {
    padding: 100px 80px;
}

.copy-container h1 {
    font-size: 100px;
    color: #DD0000;
    position:absolute;
    top: 320px;
    left: 30px;
}

.name {
    color: #fff;
    font-size: 30px;
    position:absolute;
    left: 100px;
    top: 500px;
}
.mail {
    color: #fff;
    position:absolute;
    left: 100px;
    top: 600px;
}
.pass {
    color: #fff;
    position:absolute;
    left: 100px;
    top: 700px;
}

/*メールアドレスを打ち込むところ*/
.names{
    border: 4px solid #000;
    width: 400px;
    height: 50px;
    font-size: 30px;
    background-color: #4b0082;
    color: #fff;
    position:absolute;
    left: 100px;
    top:560px;
}
.mails{
    border: 4px solid #000;
    width: 400px;
    height: 50px;
    font-size: 30px;
    background-color: #4b0082;
    color: #fff;
    position:absolute;
    left: 100px;
    top:660px;
}
.pases{
    border: 4px solid #000;
    width: 400px;
    height: 50px;
    font-size: 30px;
    background-color: #4b0082;
    color: #fff;
    position:absolute;
    left: 100px;
    top:760px;
}


.bill{
    width: 300px;
    height: 300px;
    position:absolute;
    top: 700px;
    left:600px;
}
.bills{
    width: 300px;
    height: 300px;
    position:absolute;
    top: 400px;
    left:530px;
}


button{
    width: 120px;
    height: 80px;
    font-size: 50px;
    color: #b22222;
    background-color: #4b0082;
    position:absolute;
    left: 200px;
    top:850px;

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