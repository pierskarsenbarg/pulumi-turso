# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .database import *
from .get_database import *
from .get_group import *
from .group import *
from .provider import *

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pierskarsenbarg_pulumi_turso.config as __config
    config = __config
else:
    config = _utilities.lazy_import('pierskarsenbarg_pulumi_turso.config')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "turso",
  "mod": "index",
  "fqn": "pierskarsenbarg_pulumi_turso",
  "classes": {
   "turso:index:Database": "Database",
   "turso:index:Group": "Group"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "turso",
  "token": "pulumi:providers:turso",
  "fqn": "pierskarsenbarg_pulumi_turso",
  "class": "Provider"
 }
]
"""
)
