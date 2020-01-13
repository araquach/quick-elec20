<template>
    <b-carousel
            :autoplay="true">
        <b-carousel-item v-for="(item, i) in items" :key="i">
            <figure class="image">
                <img :src="item.image">
            </figure>
            <div class="text box">
                <h1 class="title is-3 has-text-white">{{item.title}}</h1>
                <p>{{item.text}}</p>
            </div>
        </b-carousel-item>
    </b-carousel>
</template>

<script>
    export default {
        data() {
            return {
                items: [
                    {
                        title: 'Title 1',
                        text: 'Some Text',
                        image: '/dist/img/slider/plans.jpg'
                    },
                    {
                        title: 'Title 1',
                        text: 'Slide 2',
                        image: '/dist/img/slider/electrician.jpg'
                    },
                    {
                        text: 'Slide 3',
                        image: '/dist/img/slider/cctv.jpg'
                    },
                    {
                        text: 'Slide 4',
                        image: '/dist/img/slider/board.jpg'
                    },
                    {
                        text: 'Slide 5',
                        image: '/dist/img/slider/ceiling.jpg'
                    },
                    {
                        text: 'Slide 6',
                        image: '/dist/img/slider/sofa.jpg'
                    },
                    {
                        text: 'Slide 7',
                        image: '/dist/img/slider/exterior.jpg'
                    }
                ]
            }
        }
    }
</script>

<style>
    b-carousel {
        position: relative;
    }

    .text {
        position: absolute;
        top: 10%;
        left: 10%;
    }
</style>