const startVM = () => {
    vm = new Vue({
        el: '#content',
        data:{
            lyamda: 48,
            k: 0.99,
            u: 50,
            u0: 0.01,
            w: 1,
            p: Number,
            points: []
        },
        methods: {
            countP(){
                let res = 0;
                

                res = (1-this.k)*(this.w-this.k)*Math.exp(-1*(this.u-this.lyamda)*this.u0) /
                        -1*this.w*Math.expm1(-1*(this.u-this.lyamda)*this.u0)
                        
                console.log(res)
            }
        }
    })
}