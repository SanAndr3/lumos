local default = import 'default.jsonnet';

default {
  'lumos_7798-1'+: {
    config+: {
      consensus+: {
        timeout_commit: '5s',
      },
    },
  },
}
