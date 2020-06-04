var chart
var labels
var datasets = []

const startVM = () => {
    vm = new Vue({
        el: '#content',
        data:{
            lyamda: 48,
            k: 0.99,
            myu: 50,
            u0: 0.01,
            p: Number,
        },
        methods: {
            countP(){                
                let exp = Math.exp(-1*(this.myu - this.lyamda)*this.u0)

                let up = (1-this.k) * (this.w-this.k) * exp
                let down =  (this.w * exp) - 1                 
                        
                return up / down
            },
            makePoints(){
                labels = []
                let dataset = this.generateDataset()
                while(this.u0 <= 0.5){
                    labels.push(this.u0.toFixed(2))
                    dataset.data.push(this.countP() * 100)
                    this.u0 += 0.05
                    
                }
                
                datasets.push(dataset)

                this.buildChart()
            },

            generateDataset(){
                return {
                    label: `λ = ${this.lyamda},  ν = ${this.myu}`,
                    fill: false,
                    data: [],
                    borderColor: getColor(),
                    borderWidth: 2,
                }
            },

            buildChart(){
                var ctx = document.getElementById('chart');
                
                chart = new Chart(ctx, {
                    type: 'line',
                    data: {
                        labels: labels,
                        datasets: datasets
                    },
                    options: {
                        scales: {
                            yAxes: [{
                                ticks: {
                                    beginAtZero: true
                                },
                            }]
                        }
                    }
                })
            

            }

        },
        computed: {
            w(){
                return this.lyamda / this.myu
            }
        }
    })
}

function getColor(){
    return '#'+Math.floor(Math.random()*16777215).toString(16);
}