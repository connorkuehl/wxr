# wxr

Module wxr provides packages that are useful for converting a WordPress
E(x)tended RSS file into a static site.

Package wxr provides utilities for deserializing the WXR file itself
into a convenient struct.

Package markdown provides a basic way of converting an HTML parse tree
from net/html into a Markdown tree. Only a small subset of HTML to
Markdown is currently supported.

Binary cmd/wxrto utilizes this module to convert WordPress E(x)tended
RSS files into a static site. See its README.md for more information.
