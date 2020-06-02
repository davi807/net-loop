let vm
let content

/* :::::::::::::::: init ::::::::::::::: */
(() => {
    let ready = 0
    let wait

    let assets = [
        '/css/bootstrap.min.css',
        '/css/main.css',
        '/js/vue.min.js',
        '/js/components/router-params.js'
    ]

    assets.forEach(async asset => {
        insert(asset, await load(asset))
        ready += 1      
    })

    wait = setInterval(async () => {
        if(ready !== assets.length){
            return
        }
        
        clearInterval(wait)
        document.body.innerHTML = await load('/content.html')

        startVM()
    }, 50)

})()


