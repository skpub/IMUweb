<script lang="ts">
  import imu_mc from '$lib/assets/IMU_minecraft.webp'
    import { onMount } from 'svelte';
  import type { PageProps } from './$types';

  type Article = {
    id: string
    name: string
    updated: string
  }
  type ArticleGot = {
    id: string,
    name: string,
    updated: bigint
  }

  let { data }: PageProps = $props()
  onMount(() => {
    get()
  })
  let list: Article[] = $state([])
  function get() {
    data["list"]["list"].forEach((element: ArticleGot) => {
      let dateTime = new Date(Number(element["updated"]) * 1000)
      const article: Article = {
        id: element["id"],
        name: element["name"],
        updated: dateTime.toLocaleString()
      }
      list.push(article)
    });
  }
</script>

<img id="imu" src={imu_mc} alt="">
<div id="article-list-container">
  <h2 class="margin-left24">新着記事</h2>
  <div id="article-list">
    {#each list as item}
    <div class="article margin-left24">
      <p class="articleName"><b>{item["name"]}</b></p>
      <p class="articleTime">{item["updated"]}</p>
    </div>
    {/each}
  </div>
</div>
<style>
  #imu {
    width: 100%;
    display: block;
  }

  #article-list-container {
    h2 {
      margin-top: 13px;
    }
    display: flex;
    flex-flow: row;
    width: 100%;
    color: var(--white);
    background-color: var(--immoral-shadow);
    #article-list {
      flex: 1;
      display: flex;
      width: 100%;
      flex-flow: column;
      .article {
        display: flex;
        border-bottom: dotted 2px var(--white);
        .articleName {
          margin-left: 12px;
          flex: 1;
        }
        .articleTime {
          margin-right: 32px;
        }
      }
      .article:hover {
        background-color: var(--immoral-shadow-darker);
        transition: background-color 0.3s ease;
      }
    }
  }
  @media (max-width: 500px) {
    .article {
      flex-flow: column;
    }
  }
  @media (min-width: 500px) {
    .article {
      flex-flow: row;
    }
  }
</style>