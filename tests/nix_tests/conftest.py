import pytest

from .network import setup_lumos, setup_lumos_rocksdb, setup_geth


@pytest.fixture(scope="session")
def lumos(tmp_path_factory):
    path = tmp_path_factory.mktemp("lumos")
    yield from setup_lumos(path, 26650)


@pytest.fixture(scope="session")
def lumos_rocksdb(tmp_path_factory):
    path = tmp_path_factory.mktemp("lumos-rocksdb")
    yield from setup_lumos_rocksdb(path, 20650)


@pytest.fixture(scope="session")
def geth(tmp_path_factory):
    path = tmp_path_factory.mktemp("geth")
    yield from setup_geth(path, 8545)


@pytest.fixture(scope="session", params=["lumos", "lumos-ws"])
def lumos_rpc_ws(request, lumos):
    """
    run on both lumos and lumos websocket
    """
    provider = request.param
    if provider == "lumos":
        yield lumos
    elif provider == "lumos-ws":
        lumos_ws = lumos.copy()
        lumos_ws.use_websocket()
        yield lumos_ws
    else:
        raise NotImplementedError


@pytest.fixture(scope="module", params=["lumos", "lumos-ws", "lumos-rocksdb", "geth"])
def cluster(request, lumos, lumos_rocksdb, geth):
    """
    run on lumos, lumos websocket,
    lumos built with rocksdb (memIAVL + versionDB)
    and geth
    """
    provider = request.param
    if provider == "lumos":
        yield lumos
    elif provider == "lumos-ws":
        lumos_ws = lumos.copy()
        lumos_ws.use_websocket()
        yield lumos_ws
    elif provider == "geth":
        yield geth
    elif provider == "lumos-rocksdb":
        yield lumos_rocksdb
    else:
        raise NotImplementedError


@pytest.fixture(scope="module", params=["lumos", "lumos-rocksdb"])
def lumos_cluster(request, lumos, lumos_rocksdb):
    """
    run on lumos default build &
    lumos with rocksdb build and memIAVL + versionDB
    """
    provider = request.param
    if provider == "lumos":
        yield lumos
    elif provider == "lumos-rocksdb":
        yield lumos_rocksdb
    else:
        raise NotImplementedError
