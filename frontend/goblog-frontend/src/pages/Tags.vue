<template>
    <q-page padding>
        <div class="shadow-5 rounded-borders tags-container">
            <div class="q-pa-md">
                <q-chip clickable
                        v-for="tag in tagList"
                        @click="refreshBlogsWithTagId(tag.tag_id)"
                        color="deep-purple-4"
                        text-color="white"
                        icon="label"
                        :key="tag.tag_id"
                        :outline="tag.tag_id !== activeTagId">
                    {{tag.name}}
                </q-chip>
            </div>
            <q-separator />
            <div class="q-pa-md q-gutter-y-md">
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
        name: 'Tags',
        components: {
            BlogsListItem
        },
        data () {
            return {
                activeTagId: '',
                tagList: [],
                blogViewList: []
            }
        },
        methods: {
            listTagsSuccess (res) {
                console.log(res)
                if (res.data) {
                    this.tagList = res.data
                    this.activeTagId = this.tagList[0].tag_id
                    this.refreshBlogsWithTagId(this.activeTagId)
                }
            },
            listBlogsSuccess (res) {
                if (res.data) {
                    this.blogViewList = res.data
                }
            },
            refreshBlogsWithTagId (tagId) {
                axios.get('/api/blogs?tag_id=' + tagId)
                    .then(this.listBlogsSuccess)
                this.activeTagId = tagId
            }
        },
        mounted () {
            axios.get('/api/tags')
                .then(this.listTagsSuccess)
        }
    }
</script>

<style>
    .tags-container {
        margin: 0 auto;
        max-width: 1000px;
    }
</style>
