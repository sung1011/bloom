<template>
    <div class="nav-root nav flex-column nav-tabs bg-light">
        <router-link
            class="nav-link"
            :to="link.path"
            v-for="link in links"
            :key="link.name"
        >
            {{ link.title }} |
        </router-link>
    </div>
</template>

<script lang="ts" setup>
const links = [
    {
        name: 'home',
        title: 'Home',
        path: '/'
    },
    {
        name: 'about',
        title: 'About',
        path: {name: "about"}
    },
]
</script>

<style></style>

<style scoped>
.nav-root {
    width: 200px;
}
</style>