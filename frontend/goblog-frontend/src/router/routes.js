import MainLayout from 'layouts/MainLayout'
import Index from 'pages/Index'
import Blogs from 'pages/blog/Blogs'
import BlogDetail from 'pages/blog/BlogDetail'
import Types from 'pages/Types'
import Tags from 'pages/Tags'
import About from 'pages/About'
import Archives from 'pages/Archives'
import BlogsAdmin from 'pages/admin/BlogsAdmin'
import BlogEdit from 'pages/admin/BlogEdit'
import TypesAdmin from 'pages/admin/TypesAdmin'
import TagsAdmin from 'pages/admin/TagsAdmin'
import Login from 'pages/admin/Login'

const routes = [
    {
        path: '/',
        component: MainLayout,
        children: [{
            path: '/',
            component: Index
        }, {
            path: '/blogs',
            component: Blogs
        }, {
            path: '/blogs/:blog_id',
            component: BlogDetail
        }, {
            path: '/types',
            component: Types
        }, {
            path: '/tags',
            component: Tags
        }, {
            path: '/about',
            component: About
        }, {
            path: '/archives',
            component: Archives
        }, {
            path: '/admin/blogs',
            name: 'BlogsAdmin',
            component: BlogsAdmin
        }, {
            path: '/admin/blogs/edit',
            component: BlogEdit
        }, {
            path: '/admin/blogs/edit/:blog_id',
            component: BlogEdit
        }, {
            path: '/admin/types',
            component: TypesAdmin
        }, {
            path: '/admin/tags',
            component: TagsAdmin
        }, {
            path: '/admin/login',
            component: Login
        }
        ]
    },

  // Always leave this as last one,
  // but you can also remove it
  {
    path: '*',
    component: () => import('pages/Error404.vue')
  }
]

export default routes
