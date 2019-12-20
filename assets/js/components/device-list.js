(() => {
    console.log(11)
    let template = `<div>
        <span>Device list</span>
    </div>`;
    Vue.component('device-list', {
        template: template
    })
})()