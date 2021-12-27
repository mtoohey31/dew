<script>
  async function fetchHelloWorld() {
    const res = await fetch(import.meta.env.VITE_API_URL + "/");
    if (res.ok) {
      return await res.text();
    }
  }
</script>

<h1>dew</h1>

{#await fetchHelloWorld() then res}
  {res}
{/await}
