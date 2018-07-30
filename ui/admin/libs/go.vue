<template>
    <a onclick="goa(this)">
        <slot></slot>
    </a>
</template>

<script>
    window.goa = (t) => {
        var x = t.getAttribute("title");
        if (!x) x = t.innerHTML;
        document.title = x;
        var x = t.getAttribute("to");
        vm.$router.push({path: x.charAt(0) == '/' ? appUrl + x : appUrl + '/p/' + x})
    }
    export default {}
</script>

