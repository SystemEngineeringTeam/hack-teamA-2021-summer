<template>
  <div>
    <div v-for="(items, index) in itemsList" :key="index">
      <v-row justify="center" align-content="center">
        <div style="padding: 0px 50px">{{ items.name }}</div>
        <div v-for="(item, index) in items.item" :key="index">
          <div v-bind:class="item.selected ? selected : null">
            <v-col
              v-bind:class="item.selected ? selected : null"
              cols="12"
              sm="8"
              md="6"
              lg="4"
              xl="3"
            >
              <img
                :src="item.src"
                alt=""
                @click="ImageSelect(items.item, item)"
                width="150"
                height="100"
              />
            </v-col>
          </div>
        </div>
      </v-row>
    </div>
    <div class="copy-container">
      <v-btn @click="PostSetting">保存</v-btn>

      <router-link to="/main">
        <v-btn>スロットに戻る </v-btn>
      </router-link>
    </div>
  </div>
</template>
<script>
export default {
  name: "imageSelector",
  data: () => {
    return {
      selected: "selected",
      itemsList: [],
    };
  },
  methods: {
    ImageSelect(items, item) {
      items.forEach((element) => {
        element.selected = false;
      });
      item.selected = !item.selected;
    },
    async PostSetting() {
      var selected = [];
      this.itemsList.forEach((element) => {
        var select = element.item.filter((item) => item.selected)[0].src;
        select = select
          .replace("http://localhost:8080/static/img/", "")
          .replace("/", "");
        select = select.replace(/set\d/, "");
        selected.push(select);
      });
      // var selected = this.itemsList.filter(item => item.selected);
      console.log(selected.join(":"));
      try {
        var res = await this.axios.post(
          "http://localhost:8080/setting",
          {
            path: selected.join(":"),
          },
          {
            headers: {
              Authorization: "Bearer " + this.$cookie.get("token"),
            },
          }
        );
        console.log(res);
        alert("保存しました")
      } catch (e) {
        console.log(e);
      }
    },
  },
  async created() {
    var res = await this.axios.get("http://localhost:8080/images");
    var data = res.data.path;
    var item = [];
    data.forEach((paths, index) => {
      console.log(index);
      paths.forEach((path) => {
        var src = {
          src:
            "http://localhost:8080/static/img/set" + (index + 1) + "/" + path,
          selected: false,
        };
        item.push(src);
      });
      this.itemsList.push({ name: index, item: item });
      item = [];
    });
    console.log(this.itemsList);
    console.log(data);
  },
};
</script>
<style>
.selected {
  background-color: #15ff00;
}
</style>
