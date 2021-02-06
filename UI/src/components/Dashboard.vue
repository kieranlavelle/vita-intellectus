<template>
    <v-container fill-height fluid>
    </v-container>
</template>

<script>
export default {
    name: "Dashboard",
    mounted(){
        if(window.localStorage.getItem('username') == "null"){
            this.$router.push('login');
        }
    }
}
</script>