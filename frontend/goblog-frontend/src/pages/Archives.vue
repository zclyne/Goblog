<template>
    <q-page padding>
        <div class="shadow-5 archives-container q-pa-md">
            <div v-for="(blogs, key) in blogsGroupedByYear" :key="key">
                <div class="text-h3">{{key}}</div>
                <q-separator />
                <blogs-list-item
                        v-for="blog in blogs"
                        :key="blog.blog_id"
                        :blog="blog.blog"
                        :type="blog.type"
                        :tags="blog.tags"></blogs-list-item>
            </div>
        </div>
    </q-page>
</template>

<script>
    import BlogsListItem from 'components/blog/BlogsListItem'
    import axios from 'axios'
    export default {
        name: 'Archives',
        components: {
            BlogsListItem
        },
        data () {
            return {
                blogsGroupedByYear: null
            }
        },
        methods: {
            listBlogsGroupedByYearSuccess (res) {
                console.log(res)
                if (res.data) {
                    this.blogsGroupedByYear = res.data
                }
            }
        },
        mounted () {
            axios.get('/api/blogs-archive')
                .then(this.listBlogsGroupedByYearSuccess)
        }
    }
</script>

<style scoped>
    .archives-container {
        max-width: 1000px;
        margin: 0 auto;
    }
</style>
