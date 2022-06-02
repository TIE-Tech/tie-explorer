import { createRouter, createWebHistory } from 'vue-router'
import DashBoard from '../views/DashBoard'
import Block from "@/views/Block";
import Blocks from "@/views/Blocks";
import Transactions from "@/views/Transactions";
import Transaction from "@/views/Transaction";
import Tokens from "@/views/Tokens";
import Token from "@/views/Token";
import Address from "@/views/Address";
import VerifyContract from "@/views/VerifyContract";
import SourceCode from "@/views/SourceCode";
import InputJson from "@/views/InputJson";
import VyperContract from "@/views/VyperContract";

const routes = [
  {
    path: '/',
    name: 'dashboard',
    component: DashBoard
  },
  {
    path: '/blocks',
    name: 'blocks',
    component: Blocks
  },
  {
    path: '/block/:number(\\d+)',
    name: 'block',
    component: Block
  },
  {
    path: '/transactions',
    name: 'transactions',
    component: Transactions
  },
  {
    path: '/transaction/:hash',
    name: 'transaction',
    component: Transaction
  },
  {
    path: '/tokens/:type',
    name: 'tokens',
    component: Tokens
  },
  {
    path: '/token/:address',
    name: 'token',
    component: Token
  },
  {
    path: '/address/:address',
    name: 'address',
    component: Address
  },
  {
    path: '/contract_verify',
    name: 'contractVerify',
    component: VerifyContract
  },
  {
    path: '/contract_verify/source-code',
    name: 'sourceCode',
    component: SourceCode
  },
  {
    path: '/contract_verify/input-json',
    name: 'inputJson',
    component: InputJson
  },
  {
    path: '/contract_verify/vyper-contract',
    name: 'vyperContract',
    component: VyperContract
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
