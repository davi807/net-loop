let mui
let vm

function makeSidebar(){
    vm = new Vue({
        el: '#sidebar',
        data: {
            mui: mui
        },
        created(){
            document.title = this.mui.title
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
        if(window && window.document && typeof i18n === 'string' && ready !== assets.length){
            return
        }
        
        clearInterval(wait)

        document.body.innerHTML = await load('/content.html')
        mui = JSON.parse(i18n)
        makeSidebar()

    }, 50)

})()


