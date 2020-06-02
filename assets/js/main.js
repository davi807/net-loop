let mui
let vm
let content

function makeSidebar(){
    vm = new Vue({
        el: '#sidebar',
        data: {
            page: 'stat-page',
        },
        methods: {
            createVM(name){
                let path = `/js/components/${name}.js`
                let create = () => {
                    if(content && content.$destroy)
                        content.$destroy()
                    
                    content = new Vue({
                        el: '#content',
                        data: {
                            component: name
                        },
                        template: `<div id="content">
                        <component :is="component"></component>
                        </div>`
                    })
                }
                if(document.querySelector('#'+path.match(/[a-zA-Z]/g, '').join(''))){
                    create();
                } else {
                    load(path).then(res => {
                        insert(path, res)
                        create(name)    
                    })
                }

            },
            setPage(page){
                this.page = page
                this.createVM(page)
            }
        },
        created(){
            document.title = "Net-tools"
        },
        mounted(){
            this.createVM(this.page)    
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
        '/js/vue.min.js'
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
        makeSidebar()

    }, 50)

})()


