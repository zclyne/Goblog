<template>
    <q-page padding>
        <q-card class="about-card">
            <q-card-section class="q-pa-md row" horizontal>
                <div class="col-1">
                    <q-avatar size="50px" class="q-mb-sm">
                        <img src="../static/avatar.jpg">
                    </q-avatar>
                </div>
                <div class="col">
                    <div class="column">
                        <div class="col-2 text-h5">Yifan Zhang</div>
                    <div class="col text-grey-8">MSCS@Columbia University</div>
                    </div>
                </div>
            </q-card-section>
            <q-separator />
            <q-card-section>
                <q-list dense class="about-detail">
                    <q-item-label header>Experience</q-item-label>
                    <q-item>Bachelor of Engineering in Electronic Engineering@Shanghai Jiao Tong University (2016.9 ~ 2020.6)</q-item>
                    <q-item>Master of Science in Computer Science@Columbia University (Expected)</q-item>
                    <q-item>Software Developer Intern@Microsoft (2019.7 ~ 2019.9)</q-item>
                    <q-item>Software Developer Intern@VMware (2020.2 ~ Present)</q-item>
                </q-list>
            </q-card-section>
            <q-separator />
            <q-card-section>
                <q-list dense class="about-detail">
                    <q-item-label header>Tech Stack</q-item-label>
                    <q-item>Java: Spring, Spring Boot</q-item>
                    <q-item>Golang: Gin, GORM</q-item>
                    <q-item>JavaScript: Vue.js</q-item>
                </q-list>
            </q-card-section>
            <q-separator />
            <q-card-actions>
                
                <q-btn flat round icon="code">
                    <q-popup-proxy>
                        <q-banner>
                            https://github.com/zclyne
                        </q-banner>
                    </q-popup-proxy>
                </q-btn>
                <q-btn flat round icon="email">
                    <q-popup-proxy>
                        <q-banner>
                            zyfinori@gmail.com
                        </q-banner>
                    </q-popup-proxy>
                </q-btn>
            </q-card-actions>
        </q-card>
    </q-page>
</template>

<script>
    export default {
        name: 'About'
    }
</script>

<style scoped>
    .about-card {
        max-width: 1000px;
        margin: 0 auto;
    }
    .about-detail {
        font-size: 16px;
    }
</style>
