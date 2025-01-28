<script lang="ts">
  let faculties = [
    {
      name: '🎵騒音学部',
      description: 'うるさいです。',
      dean: 'まぼのすけたろ',
      link: '6o_dwir'
    },
    {
      name: '🔞実践ハラスメント学部',
      description: '終わっています。',
      dean: '佐藤海音',
      link: 'OMGR_dearinsu'
    },
    {
      name: '🎩自称マナー講師学部',
      description: '余計なマナーをクリエイトして水を差します。',
      dean: '佐藤海音(暫定)',
      link: 'OMGR_dearinsu'
    }
  ]
</script>

<div class="margin-case">
  {#each faculties as faculty}
    <h1>{faculty.name}</h1>
    <div class="margin-left24">
      <p>{faculty.description}</p>
      <p>学部長: {faculty.dean}<a href="https://x.com/{faculty.link}">(@{faculty.link})</a></p>
    </div>
  {/each}
</div>
