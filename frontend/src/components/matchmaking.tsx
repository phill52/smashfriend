import { useState } from "react";

function MatchmakingComponent() {
  const [selectedValue, setSelectedValue] = useState("");
  const [gameMode, setGameMode] = useState("Unranked");
  const [format, setFormat] = useState("");

  const handleGameModeClick = (gameMode) => {
    setGameMode(gameMode);
    setFormat("");
  };

  const handleQueueButton = () => {
    console.log(`Queue button clicked: ${gameMode} ${format}`);
  };

  const gameModeObject: Record<string, string[]> = {
    Unranked: ["Bo3", "Bo5", "Any"],
    Ranked: ["Bo3"],
  };

  const gameModeKeys = Object.keys(gameModeObject);

  return (
    <div className="flex justify-center">
      <div className="border-purple-midnight text-white-700 m-auto h-72 w-96 justify-center rounded-md border bg-[#171721] text-center font-bold text-white">
        <h1 className="pt-4 text-2xl">Find a Match</h1>
        <h2 className="text-smokey-white">Mode</h2>
        <div className="flex justify-center gap-2 px-4 py-2">
          {gameModeKeys.map((key) => (
            <button
              key={key}
              className={`mr-2 w-24 cursor-pointer rounded-xl border ${gameMode === key ? "bg-hot-pink border-neon-pink" : "border-[#7165AA] bg-[#201F2f]"} px-4 py-2 text-xs`}
              onClick={() => handleGameModeClick(key)}
            >
              {key}
            </button>
          ))}
        </div>
        <h2 className="text-smokey-white">Best of</h2>
        <div className="m-auto flex flex-col justify-center">
          <div className="my-3 flex">
            {}
            {gameModeObject[gameMode]?.map((value) => (
              <button
                key={value}
                className={`${format === value ? "bg-hot-pink border-neon-pink" : "border-[#7165AA] bg-[#201F2f]"} m-auto cursor-pointer self-start rounded-xl border px-10 py-2 text-xs`}
                onClick={() => setFormat(value)}
              >
                {value}
              </button>
            ))}
          </div>
          <button
            className="bg-pink-purple m-auto cursor-pointer self-start rounded-xl border border-[#38334E] px-10 py-2 text-xs"
            onClick={() => handleQueueButton()}
          >
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
                onChange={(event) => setSelectedValue(event?.target.value)}
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
                onChange={(event) => setSelectedValue(event.target.value)}
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
