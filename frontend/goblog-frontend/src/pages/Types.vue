<template>
    <q-page padding>
        <div class="shadow-5 rounded-borders types-container">
            <div class="q-pa-md">
                <q-chip clickable
                        v-for="type in typeList"
                        @click="refreshBlogsWithTypeId(type.type_id)"
                        color="secondary"
                        text-color="white"
                        icon="bookmark"
                        :key="type.type_id"
                        :outline="type.type_id !== activeTypeId">
                    {{type.name}}
                </q-chip>
            </div>
            <q-separator />
            <div class="q-pa-md">
                <blogs-list-item
                        v-for="item in blogViewList"
                        :key="item.blog_id"
                        :blog="item.blog"
                        :type="item.type"
                        :tags="item.tags"></blogs-list-item>
            </div>
        </div>
    </q-page>
</template>

<script>
    import BlogsListItem from 'components/blog/BlogsListItem'
    import axios from 'axios'
    export default {
        name: 'Types',
        components: {
            BlogsListItem
        },
        data () {
            return {
                activeTypeId: '',
                typeList: [],
                blogViewList: []
            }
        },
        methods: {
            listTypesSuccess (res) {
                console.log(res)
                if (res.data) {
                    this.typeList = res.data
                    this.activeTypeId = this.typeList[0].type_id
                    this.refreshBlogsWithTypeId(this.activeTypeId)
                }
            },
            listBlogsSuccess (res) {
                if (res.data) {
                    this.blogViewList = res.data
                }
            },
            refreshBlogsWithTypeId (typeId) {
                axios.get('/api/blogs?type_id=' + typeId)
                    .then(this.listBlogsSuccess)
                this.activeTypeId = typeId
            }
        },
        mounted () {
            axios.get('/api/types')
                .then(this.listTypesSuccess)
        }
    }
</script>

<style scoped>
    .types-container {
        margin: 0 auto;
        max-width: 1000px;
    }
</style>
