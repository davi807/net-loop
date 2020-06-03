const startVM = () => {
    vm = new Vue({
        el: '#content',
        data:{
            lyamda: 48,
            k: 0.99,
            myu: 50,
            u0: 0.01,
            p: Number,
            points: []
        },
        methods: {
            countP(){                
                let exp = Math.exp(-1*(this.myu - this.lyamda)*this.u0)

                let up = (1-this.k) * (this.w-this.k) * exp
                let down =  (this.w * exp) - 1                 
                        
                return up / down
            },
            makePoints(){
                this.points = []
                while(this.u0 <= 0.5){
                    this.points.push({
                        u: this.u0, 
                        p: this.countP() * 100
                    })
                    this.u0 += 0.05
                }
                
            }
        },
        computed: {
            w(){
                return this.lyamda / this.myu
            }
        }
    })
}