<template>
  <div class="main">
    <div class="copy-container">
      <h1>画像の設定</h1>
    </div>
    <selector></selector>
    <chart></chart>
  </div>
</template>

<script>
import Chart from '@/components/Chart.vue';
import Selector from '@/components/ImageSelector.vue';
export default {
  data: () => {
    return {
    };
  },
  components: {
    Chart,
    Selector
  },
  methods: {
    
  },
};
</script>
<style scoped>
body {
  font-family: "Avenir Next";
}
</style>
