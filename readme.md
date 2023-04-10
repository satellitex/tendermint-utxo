# utxo
**utxo** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Web Frontend

Ignite CLI has scaffolded a Vue.js-based web app in the `vue` directory. Run the following commands to install dependencies and start the app:

```
cd vue
npm install
npm run serve
```

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/username/utxo@latest! | sudo bash
```
`username/utxo` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)


## Default config:
```
Keys #1: (618741408250DF80E5BC33999AE29A7E09FEBBD4, 618741408250DF80E5BC33999AE29A7E09FEBBD4, e3d9f57f0a5082b5e6d95d86d801f4581a026529295fa0a6fc224886b8c39bf8)
Keys #2: (D864BB8367F7963FE9BA6BC452533C15F1F415FE, D864BB8367F7963FE9BA6BC452533C15F1F415FE, 4e816874d6e1473811e8eb15f70159832307b2bf532b58eb02a43348b51dce83)
Keys #3: (CC14092F505D1284C8EF287FC97E3E0572591B2D, CC14092F505D1284C8EF287FC97E3E0572591B2D, 7a3a888d20bf2192524e88059ac6e6b0cbcf435bfa7c32187992d238409c298d)
Keys #4: (DE1C8C1362C55E076CBEC3855BDC47D8A95CFEC8, DE1C8C1362C55E076CBEC3855BDC47D8A95CFEC8, 53c4236dc46bf5c8ae8e82e7e54d4a0384c3bd0dbcca1c28a9ee43b5e4538a3d)
Keys #5: (2E5E6C36ABDBD3379001E346A834EE56E9E3CD37, 2E5E6C36ABDBD3379001E346A834EE56E9E3CD37, d5d52b1a97f5e179632d472fd87ae752494c8d43ceb7b2f90098d0add3011ed7)
```

## TODO
- [x] Balance の残高参照処理追加
- [x] Genesis の適切な作成
- [ ] CLI の手入れ
- [ ] 正常系ストーリー作成