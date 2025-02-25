<script lang="ts">
  import { PUBLIC_BACKEND_ADDR, PUBLIC_BACKEND_PORT } from '$env/static/public'
  import imu_mc from '$lib/assets/IMU_minecraft.webp'
  import { onMount } from 'svelte'
  import type { PageProps } from './$types'
  import { marked } from 'marked'

  type ArticleHead = {
    id: string
    title: string
    name: string,
    studentID: string,
    updated: string
  }
  type ArticleHeadGot = {
    id: string,
    title: string,
    name: string,
    studentID: string,
    updated: bigint
  }

  let { data }: PageProps = $props()
  onMount(() => {
    get()
  })
  let list: ArticleHead[] = $state([])
  function get() {
    if (data["list"]["list"] === undefined) return
    data["list"]["list"].forEach((element: ArticleHeadGot) => {
      console.log(element)
      let dateTime = new Date(Number(element.updated) * 1000)
      const article: ArticleHead = {
        id: element.id,
        title: element.title,
        name: element.name,
        studentID: element.studentID,
        updated: dateTime.toLocaleString()
      }
      list.unshift(article)
    })
  }

  let images: Record<string, string> = $state({})

  type Article = {
    title: string,
    content: string,
  }
  let article: Article | undefined = $state(undefined)
  let parsedArticle = $state("")

  function arrayBufferToDataURL(binary: string, mimeType: string): string {
    const buffer = new Uint8Array(binary.length)
    for (let i = 0; i < binary.length; i++) {
      buffer[i] = binary.charCodeAt(i)
    }
    const blob = new Blob([buffer], { type: mimeType })
    return URL.createObjectURL(blob)
  }

  const getArticle = async (id: string) => {
    const res = await fetch(PUBLIC_BACKEND_ADDR + ':' + PUBLIC_BACKEND_PORT + '/api/article/get/' + id)
    const data = await res.json()
    const article_: Article = {
      title: data["articleTitle"],
      content: data["content"]
    }
    article = article_
    if (data["image"] !== undefined) {
      data["image"].forEach((element: Record<string, any>) => {
        const mimeType = element["name"].includes(".png") ? "image/png" : "image/jpeg"
        const binary = atob(element["content"])
        images[element["name"]] = arrayBufferToDataURL(binary, mimeType)
      });
    }
    parsedArticle = await processMarkdown(article.content)
  }

  async function processMarkdown(md: string): Promise<string> {
    const renderer = new marked.Renderer()
    renderer.image = ({ href, title, text }) => {
      const dataURL = images[href.slice(2)] || images[href] || href
      //                          ↑ これ、 ./img.png みたいなのを想定している。
      //                            将来的にはちゃんと正規表現で書くか、CMS側を作ってからどうにかする。
      return `<img src="${dataURL}" alt="${text}" title="${title || ''}"/>`
    }
    return await marked(md, { renderer })
  }
</script>

<img id="imu" src={imu_mc} alt="">
<div id="article-list-container">
  <h2 class="margin-left24">新着記事</h2>
  <div id="article-list">
    {#each list as item}
    <div class="articleSelection margin-left24" onclick={() => getArticle(item["id"])}>
      <p class="articleTitle"><b>{item.title}</b></p>
      <div>
        <p class="articleDesc">{item.name}@{item.studentID} {item.updated}</p>
      </div>
    </div>
    {/each}
  </div>
</div>
<div id="article" class="margin-case">
  {@html parsedArticle}
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
      .articleSelection {
        display: flex;
        border-bottom: dotted 2px var(--white);
        .articleTitle {
          margin-left: 12px;
          flex: 1;
        }
        .articleDesc {
          margin-right: 32px;
        }
      }
      .articleSelection:hover {
        background-color: var(--immoral-shadow-darker);
        transition: background-color 0.3s ease;
      }
    }
  }
  #article {
    box-sizing: border-box;
    overflow-x: hidden;
  }
  :global(#article img) {
    max-width: 100%;
    height: auto;
    display: block;
    object-fit: contain;
  }
  @media (max-width: 500px) {
    .articleSelection {
      flex-flow: column;
    }
  }
  @media (min-width: 500px) {
    .articleSelection {
      flex-flow: row;
    }
  }
</style>