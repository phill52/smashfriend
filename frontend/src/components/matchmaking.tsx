import { useState } from "react";

function MatchmakingComponent() {
  const [selectedValue, setSelectedValue] = useState("");

  const handleChange = (event) => {
    console.log("selected ", event.target.value);
    setSelectedValue(event.target.value);
  };

  return (
    <div className="flex h-120 justify-center">
      <div className="text-white-700 m-auto h-72 w-96 justify-center rounded-md bg-[#171721] text-center font-bold text-white">
        <h1 className="pt-4 text-2xl">Find a Match</h1>
        <h2 className="text-smokey-white">Mode</h2>
        <div className="flex justify-center gap-2 px-4 py-2">
          <button className="mr-2 w-24 rounded-xl border border-[#7165AA] bg-[#201F2f] px-4 py-2 text-xs">
            Unranked
          </button>
          <button className="bg-hot-pink border-neon-pink w-24 rounded-xl border text-xs">
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
        <div className="mt-2 flex justify-center gap-2 text-xs">
          <p className="">Hide Username</p>
          <div className="flex gap-2">
            <label>
              <input
                className="mr-1"
                id="on"
                name="on"
                value="on"
                type="radio"
                checked={selectedValue === "on"}
                onChange={handleChange}
              />
              On
            </label>
            <label>
              <input
                className="mr-1"
                type="radio"
                id="off"
                name="off"
                value="off"
                checked={selectedValue === "off"}
                onChange={handleChange}
              />
              Off
            </label>
          </div>
        </div>
      </div>
    </div>
  );
}

export default MatchmakingComponent;
