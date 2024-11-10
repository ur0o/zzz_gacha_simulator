import { Routes, Route } from 'react-router-dom'
import { Home } from "./pages/Home.tsx"
import { NotFound } from "./pages/NotFound"
import { css } from '@emotion/react'
import { Gacha } from './pages/Gacha.tsx';

export default function App() {
  const ContainerStyle = css`
    width: 1080px;
    margin: 0 auto;
    padding: 16px;
  `;

  return <>
    <div css={ContainerStyle}>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/gacha/">
          <Route path=":gachaType/:id" element={<Gacha />}/>
        </Route>
        <Route path="" element={<NotFound />} />
      </Routes>
    </div>
  </>
}
