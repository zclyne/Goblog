<template>
    <q-page padding>
        <blogs-list-item
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
