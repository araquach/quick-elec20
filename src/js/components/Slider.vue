<template>
  <b-carousel
      :autoplay="true"
      :interval="5000"
      :pause-hover="false"
      :pause-info="false">
    <b-carousel-item v-for="(item, i) in items" :key="i">
      <div :class="item.image" class="hero is-medium">
        <div class="hero-body">
          <div class="box">
            <h1 class="title is-4 has-text-white">{{ item.title }}</h1>
            <p class="is-size-5">{{ item.text }}</p>
            <a :href="item.link" class="has-text-white is-size-5">{{ item.linkText }}</a>
          </div>
        </div>
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
          title: 'Make your home a SMART home',
          text: 'Smart home features are not as expensive as you might think',
          image: 'smarthome',
          link: '/smarthome',
          linkText: 'Find Out More >',
          side: 'is-left'
        },
        {
          title: 'Expert residential and commercial electricians',
          text: 'Years of experience, friendly & reliable service',
          image: 'plans',
          link: '/about',
          linkText: 'Find Out More >',
          side: 'is-left'
        },
        {
          title: 'Full and partial rewires',
          text: 'We offer a full range of services to the homeowner from fitting a couple of sockets or light fitting to a complete rewire',
          image: 'electrician',
          link: '/about',
          linkText: 'More Info >',
          side: 'is-right'
        },
        {
          title: 'Burglar Alarm & CCTV installation',
          text: 'Ensure your home or business are fully secured. Quick-Elec will install burglar alarms and CCTV systems to suit your budget and your needs',
          image: 'cctv',
          link: '/about',
          linkText: 'Find Out More >',
          side: 'is-left'
        },
        {
          title: 'Consumer unit upgrades',
          text: 'All household consumer units should be RCD protected. An RCD, or residual current device, is a life-saving device which is designed to prevent you from getting a fatal electric shock',
          image: 'board',
          link: '/about',
          linkText: 'Find Out More >',
          side: 'is-right'
        },
        {
          title: 'Lighting updates',
          text: 'Upgrade the lighting in your home with down lights, led lighting and security lighting. Whatever your lighting needs Quick-Elec will deliver',
          image: 'ceiling',
          link: '/domestic',
          linkText: 'Click for more info >',
          side: 'is-left'
        },
        {
          title: 'Energy saving upgrades',
          text: 'Upgrade your home or business with Quick-Elec and save upto 30% on your energy bills',
          image: 'sofa',
          link: '/about',
          linkText: 'More Here >',
          side: 'is-right'
        },
        {
          title: 'Add lighting to your outdoor space',
          text: 'Add a touch of class to your garden or patio area',
          image: 'exterior',
          link: '/exterior',
          linkText: 'Find Out More >',
          side: 'is-left'
        },
        // {
        //   title: '',
        //   text: '',
        //   image: 'open',
        //   link: '/about',
        //   linkText: '',
        //   side: 'is-left'
        // },
      ]
    }
  }
}
</script>
<style lang="scss" scoped>
.carousel-item:last-of-type {
  .box {
    opacity: 0;
  }
}

.hero {
  background-size: cover;
}

.open {
  background-image: url("/dist/img/slider/open.jpg");
  background-position: center center;
}

.plans {
  background-image: url("/dist/img/slider/plans.jpg");
}

.electrician {
  background-image: url("/dist/img/slider/electrician.jpg");
}

.cctv {
  background-image: url("/dist/img/slider/cctv.jpg");
}

.board {
  background-image: url("/dist/img/slider/board.jpg");
}

.ceiling {
  background-image: url("/dist/img/slider/ceiling.jpg");
}

.sofa {
  background-image: url("/dist/img/slider/sofa.jpg");
}

.exterior {
  background-image: url("/dist/img/slider/exterior.jpg");
}

.smarthome {
  background-image: url("/dist/img/slider/smarthome.jpg");
  background-position: center center;
}
</style>