<script lang="ts">
  import type { PageProps } from './$types'
  let { data }: PageProps = $props()
</script>

<div class="margin-case">
  <h1>マイクラサーバ接続情報</h1>
  <div class="margin-left24">
    <div class="mctable">
      <span>バージョン</span>
      <span>{data.version}</span>
    </div>
    <div class="mctable">
      <span>JE アドレス:ポート </span>
      <span>{data.JEAddr}</span>
    </div>
    <div class="mctable">
      <span>BE アドレス:ポート </span>
      <span>{data.BEAddr}</span>
    </div>
  </div>
</div>

<style>
  .mctable {
    display: flex;
    font-family: monospace, serif;
  }
  .mctable span:nth-child(odd) {
    width: 180px;
  }
</style>
