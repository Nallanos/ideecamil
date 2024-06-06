<script lang="ts">
  const API_URL = "http://localhost:8080";

  type Result = {
    X: number;
    Y: number;
  };

  let result: Result;
  let finalString: string;
  let arrivalX = "";
  let arrivalY = "";

  let vecteurX = "";
  let vecteurY = "";

  // const updateArrivalVal = (event: Event, val: string) => {
  //   let inputElement = event.target as HTMLInputElement;
  //   arrivalPoint = inputElement.value;
  // };

  async function getResult() {
    const finalInput = {
      vecteur: {
        x: vecteurX,
        y: vecteurY,
      },
      arrivalPoint: {
        x: arrivalX,
        y: arrivalY,
      },
    };
    const req = new Request(API_URL, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(finalInput),
    });

    try {
      const res = await fetch(req);
      const data = await res.json();
      console.log(finalInput);
      result = data;
      finalString = `A=(x:${result.X}, y:${result.Y})`;
      console.log(result, finalString);
    } catch (err: any) {
      console.log({ error: err.message });
    }
  }
</script>

<main>
  <form
    on:submit={async () => {
      await getResult();
    }}
  >
    <div class="first">
      <label for="">Vecteur X</label>
      <input type="number" bind:value={vecteurX} />

      <label for="">Vecteur Y</label>
      <input type="number" bind:value={vecteurY} />
    </div>
    <div>
      <label for="">Arrival Point X</label>
      <input type="number" bind:value={arrivalX} />
      <label for="">Arrival Point Y</label>
      <input type="number" bind:value={arrivalY} />
    </div>
    <button type="submit" class="button">Submit</button>
  </form>
  <div>
    <h1>Result</h1>
    <p>{finalString}</p>
  </div>
</main>

<style>
  .button {
    width: 80px;
  }
  .first {
    margin-bottom: 5px;
  }
</style>
