local config = import 'default.jsonnet';

config {
  'lumos_7798-1'+: {
    config+: {
      storage: {
        discard_abci_responses: true,
      },
    },
  },
}
