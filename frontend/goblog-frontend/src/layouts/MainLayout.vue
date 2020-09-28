<template>
    <q-layout view="lHh Lpr lFf">
        <q-header elevated>
            <q-toolbar>
                <q-btn
                        flat
                        dense
                        round
                        icon="menu"
                        aria-label="Menu"
                        @click="leftDrawerOpen = !leftDrawerOpen"
                />

                <q-toolbar-title>
                    Goblog
                </q-toolbar-title>

                <q-btn flat label="ADMIN" to="/admin/login" />
            </q-toolbar>
        </q-header>

        <q-drawer
                v-model="leftDrawerOpen"
                show-if-above
                bordered
                content-class="bg-grey-1"
        >
            <q-list style="height: calc(100% - 150px); margin-top: 150px; border-right: 1px solid #ddd">
                <EssentialLink
                        v-for="item in essentialLinks"
                        :key="item.title"
                        v-bind="item"
                />
            </q-list>
            <q-img class="absolute-top" src="https://cdn.quasar.dev/img/material.png" style="height: 150px">
                <div class="absolute-bottom bg-transparent">
                    <q-avatar size="56px" class="q-mb-sm">
                        <img src="../static/avatar.jpg">
                    </q-avatar>
                    <div class="text-weight-bold">Yifan Zhang</div>
                    <div>MSCS@Columbia University</div>
                </div>
            </q-img>
        </q-drawer>

        <q-page-container>
            <router-view />
        </q-page-container>
    </q-layout>
</template>

<script>
    import EssentialLink from 'components/EssentialLink.vue'

    const linksData = [
        {
            title: 'Blogs',
            icon: 'article',
            path: '/blogs'
        },
        {
            title: 'Types',
            icon: 'bookmark',
            path: '/types'
        },
        {
            title: 'Tags',
            icon: 'label',
            path: '/tags'
        },
        {
            title: 'Archives',
            icon: 'archive',
            path: '/archives'
        },
        {
            title: 'About Me',
            icon: 'account_circle',
            path: '/about'
        }
    ]

    export default {
        name: 'MainLayout',
        components: { EssentialLink },
        data () {
            return {
                leftDrawerOpen: true,
                essentialLinks: linksData
            }
        }
    }
</script>
