import { css } from "@emotion/react";
import { Link } from "react-router-dom";

export function Header() {
  const HeaderStyle = css`
    padding: 16px 8px;
    background-color: #333;
    color: white;
  `

  const HeaderLinkStyle = css`
    color: white;
    font-size: 1.6rem;
  `

  return <div css={HeaderStyle}>
    <Link css={HeaderLinkStyle} to="/">Header</Link>
  </div>
}