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
                <input v-model="name"/>
                <p class="mail">メールアドレス</p>
                <input v-model="mail"/>
                <p class="pass">パスワード</p>
                <input v-model="password"/>
                
            </div>
            <img class="loanshark" src="https://4.bp.blogspot.com/-VA_IOO7AQy4/VeJ-lBs7uwI/AAAAAAAAxH8/36QFYd58nRk/s800/money_yami_soshiki.png">
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
    font-family: "Avenir Next";
    background-color: #000000;
}

li {
    list-style: none;
}

/*headerの設定、背景の色*/
.header {
    background-color: #b22222;    color: #4b0082;
    height: 300px;
}

/*heder、チーム名とかロゴ的な何かの大きさ、位置*/
.header-logo {
    float: left;
    font-size: 145px;
    padding: 20px 40px;
}

/*headerおわり*/
/*mainの設定*/
.main {
    padding: 100px 80px;
}

.copy-container h1 {
    font-size: 100px;
    color: #800000;
    position:relative;
    left: 10px;
}

.name {
    color: #fff;
    position:relative;
    left: 30px;
    bottom: 40px;
}
.mail {
    color: #fff;
    position:relative;
    left: 80px;
    bottom: 40px;
}
.pass {
    color: #fff;
    position:relative;
    left: 50px;
    bottom: 40px;
}

/*メールアドレスを打ち込むところ*/
input{
    border: 4px solid #000;
    width: 400px;
    height: 30px;
    background-color: #4b0082;
    color: #fff;
    position:relative;
    left: 150px;
    bottom: 40px;
}

.loanshark{
    width:400px;
    position:relative;
    left: 95px;
    top: 50px;
}

button{
    width: 120px;
    height: 80px;
    font-size: 50px;
    color: #b22222;
    background-color: #4b0082;
    bottom: 30px;
    left: 0px;

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