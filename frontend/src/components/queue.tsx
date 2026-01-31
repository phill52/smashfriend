function QueueComponent({ gameFormat: string, gameMode: string }) {
  return (
    <div className="flex justify-center">
      <div
        className={`border-purple-midnight text-white-700 bg-deep-shadow m-auto h-72 w-96 flex flex-col justify-center items-center rounded-md border text-center font-bold text-white`}
      >
        <h1 className="text-3xl p-6">Finding Opponents...</h1>
        <div className="">
          <div className="bg-shadow-violet border border-twilight-purple rounded-xl w-40 h-18 flex items-center justify-center">
            <p className="text-3xl">0:00</p>
          </div>
        </div>
        <p className="text-smokey-white text-xs p-4">
          Mode:{" "}
          <span className="text-white">
            {gameMode}/{gameFormat}
          </span>
        </p>
        <p className="text-smokey-white text-sm">
          Estimated time: <span className="text-white">0s</span>
        </p>
      </div>
    </div>
  );
}

export default QueueComponent;
