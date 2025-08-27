function MatchmakingComponent() {
  return (
    <div className="flex h-120 w-144 justify-center rounded-md border border-[#38334E] bg-linear-to-b from-[#312843] via-[#252034] to-[#15151F]">
      <div className="text-white-700 m-auto flex h-72 w-96 justify-center rounded-md bg-[#171721] text-center font-bold text-white">
        <div>
          <div>
            <h1 className="pt-4 text-2xl">Find a Match</h1>
          </div>
          <h2 className="text-smokey-white">Mode</h2>
          <div className="flex flex-row justify-between p-1">
            <button className="rounded-xl border border-[#7165AA] bg-[#201F2f] p-2 text-xs">
              Unranked
            </button>
            <button className="bg-hot-pink border-neon-pink rounded-xl border p-2 text-xs">
              Ranked
            </button>
          </div>
          <h2 className="text-smokey-white">Best of</h2>
          <div className="m-auto flex flex-col justify-center">
            <div className="my-3">
              <button className="bg-hot-pink border-neon-pink m-auto self-start rounded-xl border px-10 py-2 text-xs">
                Bo3
              </button>
            </div>
            <button className="bg-pink-purple m-auto self-start rounded-xl border border-[#38334E] px-10 py-2 text-xs">
              Queue
            </button>
          </div>
          <div className="mt-2 flex text-xs">
            <p>Hide Username</p>
            <div>
              <label>
                <input id="On" name="On" value="On" type="radio" />
                On
              </label>
            </div>
            <div className="accent-purple">
              <label>
                <input type="radio" id="Off" name="Off" value="Off" />
                Off
              </label>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default MatchmakingComponent;
