<template>
  <div class="main">
    <div class="copy-container">
      <h1>画像の設定</h1>
    </div>
    <div class="contents">
      <div v-for="(items, index) in itemsList" :key="index">
        <v-row justify="center" align-content="center">
          {{  items.name }}
          <div v-for="(item, index) in items.item" :key="index">
            <div v-bind:class="item.selected? selected : null ">
              <v-col v-bind:class="item.selected? selected : null " cols="12" sm="8" md="6" lg="4" xl="3">
                <img :src="item.src" alt="" @click="ImageSelect(items.item, item)" width="150" height="100">
              </v-col>
            </div>
          </div>
        </v-row>
        <!-- <hr> -->
      </div>
    </div>
    <div class="copy-container">
      <v-btn @click="PostSetting">保存</v-btn>
    </div>
  </div>
</template>

<script>
export default {
  data: () => {
    return {
      selected: 'selected',
      itemsList: []
    };
  },
  methods: {
    ImageSelect(items, item) {
      items.forEach(element => {
        element.selected = false;
      });
      item.selected = !item.selected
    },
    PostSetting() {
      var selected = [];
      this.itemsList.forEach(element => {
        var select = element.item.filter(item => item.selected)[0].src;
        select = select.replace("http://localhost:8080/static/img/", "").replace("/", "");
        select = select.replace(/set\d/, "");
        selected.push(select);
      });
      // var selected = this.itemsList.filter(item => item.selected);
      console.log(selected.join(':'));
      this.axios.post('http://localhost:8080/setting', {
        path: selected.join(':'),
        email: "test@test" //emailに置き換えてください
      }
      ).then(function (response) {
        console.log(response);
      }).catch(function (error) {
        console.log(error);
      });
    },
  },
  async created() {
    var res = await this.axios.get("http://localhost:8080/images");
    var data = res.data.path;
    var item = [];
    data.forEach((paths, index) => {
      console.log(index);
      paths.forEach(path => {
        var src = {src: 'http://localhost:8080/static/img/set' + (index + 1) + "/" + path, selected: false};
        item.push(src);
      });
      this.itemsList.push({name: index, item: item});
      item = [];
    });
    // console.log(this.itemsList);
    // console.log(data);
  },
};
</script>
<style scoped>
body {
  font-family: "Avenir Next";
}

li {
  list-style: none;
}

/*headerの設定、背景の色*/
.header {
  background-color: #26d0c9;
  color: #fff;
  height: 90px;
}

/*heder、チーム名とかロゴ的な何かの大きさ、位置*/
.header-logo {
  float: left;
  font-size: 36px;
  padding: 20px 40px;
}

/*heder、項目の大きさ、位置*/
.header-list li {
  float: left;
  padding: 33px 20px;
}
/*headerおわり*/
/*mainの設定*/
.main {
  padding: 100px 80px;
}

.copy-container h1 {
  font-size: 100px;
}

/*メールアドレスを打ち込むところ*/
input {
  border: 3px solid #000;
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

.contents-item p {
  font-size: 24px;
  margin-top: 30px;
}

/*フッターの設定*/
.footer {
  background-color: #2f5167;
  color: #fff;
  height: 120px;
  padding: 40px;
}

.footer-logo {
  float: left;
  font-size: 32px;
}

.footer-list {
  float: right;
}

.footer-list li {
  padding-bottom: 20px;
}

.selected {
  background-color: #15ff00;
}
</style>
