<template>
    <q-page padding class="q-gutter-y-md">
        <blogs-list-item
                class="shadow-5"
                v-for="item in blogViewList"
                :key="item.blog_id"
                :blog="item.blog"
                :type="item.type"
                :tags="item.tags"></blogs-list-item>
    </q-page>
</template>

<script>
    import BlogsListItem from 'components/blog/BlogsListItem'
    import axios from 'axios'
    export default {
        name: 'Blogs',
        components: {
            BlogsListItem
        },
        data () {
            return {
                blogViewList: []
            }
        },
        methods: {
            listBlogsSuccess (res) {
                console.log(res)
                if (res.data) {
                    this.blogViewList = res.data
                }
            }
        },
        mounted () {
            axios.get('/api/blogs')
                .then(this.listBlogsSuccess)
        }
    }
</script>
