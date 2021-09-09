<script>
import { Line } from 'vue-chartjs';

export default {
  extends: Line,
  name: 'LineChart',
  props: ['user'],
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
    var data = await this.axios.get('http://localhost:8080/chartdata', {
      headers: {
        email: this.user
      },
    });
    data.data.data.forEach(element => {
      this.chartdata.labels.push(element.spin);
      this.chartdata.datasets[0].data.push(element.coin);
    });
    this.renderChart(this.chartdata, this.options);
  },
}
</script>