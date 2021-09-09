<script>
import { Line } from 'vue-chartjs';

export default {
  extends: Line,
  name: 'LineChart',
  data:() => ({
    chartdata: {
      labels: [],
      datasets: [
        {
          label: 'Coin',
          data: [],
        }
      ]
    },
    options: {
      responsive: true,
      maintainAspectRatio: false
    }
  }),
  async mounted (){
    try {
      var data = await this.axios.get('http://localhost:8080/chartdata', {
        headers: {
          Authorization: 'Bearer ' + this.$cookie.get("token")
        },
      });    
      data.data.data.forEach(element => {
        this.chartdata.labels.push(element.spin);
        this.chartdata.datasets[0].data.push(element.coin);
      });
      console.log(data)
      this.renderChart(this.chartdata, this.options);
    } catch (error) {
      console.log(error);
    }
  },
}
</script>