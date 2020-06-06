var chart
var labels = []
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
                while(this.u0 <= 0.4){
                    labels.push(this.u0.toFixed(2))
                    dataset.data.push(this.countP() * 100)
                    this.u0 += 0.03
                    
                }
                
                datasets.push(dataset)
                this.u0 = 0.01

                this.buildChart()
            },

            generateDataset(){
                let color = getColor()
                return {
                    label: `λ = ${this.lyamda},  ν = ${this.myu}`,
                    fill: false,
                    data: [],
                    borderColor: color,
                    backgroundColor: color,

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
                        responsive: true,
                        scales: {
                            yAxes: [{
                                ticks: {
                                    beginAtZero: true
                                },
                            }]
                        }
                    }
                })
            

            },

            reset(){
                this.lyamda = 48
                this.k = 0.99
                this.myu = 50
                this.u0 = 0.01
                this.p = Number
                
                if(chart){
                    chart.destroy()
                }
 
                labels = []
                datasets = []
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
    let colors = [
        'ff4000',
        '00ff00',
        '0080ff',
        '8000ff',
        'ff0040',
        'ffbf00',
        '4d79ff'
    ]
    return '#'+colors[Math.floor(Math.random() * colors.length)]
}