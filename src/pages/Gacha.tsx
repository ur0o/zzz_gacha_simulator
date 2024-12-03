import { css } from "@emotion/react";
import { BaseUrl } from "../utils/api_config";
import axios, { AxiosRequestConfig, AxiosResponse } from "axios";
import { useEffect, useState } from "react";
import { useParams } from "react-router";

export declare namespace Gacha {
  type Props = {
    gachaType: "character" | "weapon";
    id: number;
  }

  type Info = {
    id: number;
    name: string;
    puS_name: string;
    puA_0_name: string;
    puA_1_name: string;
    start_date: string;
    end_date: string;
  };
}

type Result = {
  name: string;
  rank: string;
}

type State = {
  offsetS: number;
  offsetA: number;
  fixedS: boolean;
  fixedA: boolean;
}


export function Gacha() {
  const [gachaInfo, setGachaInfo] = useState<Gacha.Info>();
  const [result, setResult] = useState<Result[]>([]);
  const [history, setHistory] = useState<Result[]>([]);
  const [puCount, setPuCount] = useState<Record<string, number>>({});
  const [gachaState, setGachaState] = useState<State>({
    offsetS: 0,
    offsetA: 0,
    fixedS: false,
    fixedA: false,
  });
  const { gachaType, id } = useParams();



  useEffect(() => {
    const options: AxiosRequestConfig = {
      baseURL: BaseUrl,
      url: `/gacha/${gachaType}/${id}`,
      method: "get",
    }
    axios(options).then((res: AxiosResponse<Gacha.Info>) => {
      const { data } = res;
      setGachaInfo(data);
    })
  }, [])

  const puNames = function() {
    return gachaInfo ?
      {
        s: gachaInfo.puS_name,
        a0: gachaInfo.puA_0_name,
        a1: gachaInfo.puA_1_name
      }
      : { s: "", a0: "", a1: "" };
  }();

  function draw() {
    const options: AxiosRequestConfig = {
      baseURL: BaseUrl,
      url: `/gacha/${gachaType}/${id}/draw`,
      method: "get",
      params: gachaState,
    }
    axios(options).then((res: AxiosResponse<Result[]>) => {
      const { data } = res;
      setResult(data);
      setHistory(result.toReversed().concat(history).slice(0, 90));
      calcState(data);
      calcPuCount(data);
    })
  }

  function calcState(data: Result[]) {
    const reversedData = data.toReversed();
    const indexA = reversedData.findIndex(el => {
      return el.rank == "A" || el.rank == "S"
    });
    const indexS = reversedData.findIndex(el => {
      return el.rank == "S"
    });

    const state = {
      offsetS: indexS === -1 ? gachaState.offsetS + 10 : indexS,
      offsetA: indexA,
      fixedS: false,
      fixedA: false,
    }
    setGachaState(state);
  }

  function calcPuCount(data: Result[]) {
    const newCount = {...puCount};
    newCount.s += data.filter((d) => d.name == puNames.s).length;
    newCount.a0 += data.filter((d) => d.name == puNames.a0).length;
    newCount.a1 += data.filter((d) => d.name == puNames.a1).length;
    setPuCount(newCount);
  }

  const containerStyle = css`
    text-align: center;
  `;

  const mainContainerStyle = css`
    display: flex;

    .content-left {
      width: 50%;

    }

    .content-right {
      width: 50%;
    }
  `;

  return <div css={containerStyle}>
    <h2>{gachaInfo ? gachaInfo.name : ""}</h2>
    <button onClick={() => draw()}>ガチャる</button>
    <div css={mainContainerStyle}>
      <div className="content-left">
        <div className="puCount">
          {Object.keys(puNames).map((p) => {
            return <div key={p}>{puNames[p]}: {puCount[p]}回</div>
          })}
        </div>
      </div>
      <div className="content-right"><ResultTable result={result} /></div>
    </div>
    <HistoryTable history={history} />
  </div>
}

function ResultTable({result}: {result: Result[]}){
  const resultStyle = css`
    margin: 0 auto;
    padding-bottom: 8px;
  `

  return <>
  { result.length == 0
    ? <div>なし</div>
    : <table css={resultStyle}>
      <thead>
        <tr>
          <th>レアリティ</th>
          <th>名前</th>
        </tr>
      </thead>
      <tbody>
        {result.map((res, i) => {
          return <tr key={i}>
            <td>{res.rank}</td>
            <td>{res.name}</td>
          </tr>
        })}
      </tbody>
    </table>
  }
  </>
}

function HistoryTable({history}: {history: Result[]}) {
  const columnsNum = Math.ceil(history.length / 5);
  const rowsNum = Math.min(history.length, 5);

  const historyTitleStyle = css`
    font-weight: bold;
  `;
  const historyStyle = css`
    display: flex;
    width: 100%;
    overflow-x: scroll;

    border: 1px solid gray;
    padding: 8px;

    .column {
      width: calc(2rem + 160px);
      white-space: nowrap;

      .unit {
        display: flex;
        align-items: center;
        line-height: 1.4rem;
        vertical-align: baseline;

        .index {
          display: inline-block;
          width: 2rem;
          text-align: right;;
        }

        .name {
          display: inline-block;
          width: 160px;

          text-align: center;
          white-space: nowrap;
          overflow: hidden;
          text-overflow: ellipsis;
        }
      }
    }
  `;

  return <>
    <div css={historyTitleStyle}>ガチャ履歴(90件まで)</div>
    { history.length == 0
      ?
      <div>履歴なし</div>
      :
      <div css={historyStyle}>
        { [...Array(columnsNum)].map((_, i) => {
          return <div className="column" key={i}>
            { [...Array(rowsNum)].map((_, j) => {
              const index = j + i * rowsNum;
              return <div className="unit" key={`${i}_${j}`}>
                <div className="index">{index + 1}</div>
                <div className="name">{history[index].name}</div>
              </div>
            })}
          </div>
        })}
      </div>
    }
  </>
}