let mui
let vm

function makeInstance(){
    mui = JSON.parse(i18n)
    vm = new Vue({
        el: '[data-container]',
        data: {
            mui: mui
        }
    })
}

(async () => {
    [
        '/css/bootstrap.min.css',
        '/css/main.css',
        '/js/font-awesome.js',
        '/js/vue.min.js'
    ].forEach(async asset => {
        insert(asset, await load(asset))         
    })

    document.body.innerHTML = await load('/content.html')    
    
    let wait = setInterval(() => {
        if(typeof Vue !== 'undefined' && typeof i18n !== 'undefined'){
            clearInterval(wait)
            makeInstance()
        }
    }, 30)

})()


