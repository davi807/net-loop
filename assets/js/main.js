let mui
let vm
let content

function makeSidebar(){
    vm = new Vue({
        el: '#sidebar',
        data: {
            mui: mui
        },
        created(){
            document.title = this.mui.title
        },
        mounted(){
            load('/js/components/home.js').then(res => {
                insert('/js/components/home.js', res)
                content = new Vue({
                    el: '#content',
                    data: {
                        component: 'start-page'
                    },
                    template: `<div id="content">
                      <component :is="component"></component>
                    </div>`
                })
            })
        }
    }) 
}

/* :::::::::::::::: init ::::::::::::::: */

(() => {
    let ready = 0
    let wait

    let assets = [
        '/css/bootstrap.min.css',
        '/css/main.css',
        '/js/font-awesome.js',
        '/js/vue.min.js'
    ]

    assets.forEach(async asset => {
        insert(asset, await load(asset))
        ready += 1      
    })

    wait = setInterval(async () => {
        if( typeof(i18n) !== 'string' || ready !== assets.length){
            return
        }
        
        clearInterval(wait)

        document.body.innerHTML = await load('/content.html')
        mui = JSON.parse(i18n)
        makeSidebar()

    }, 50)

})()


