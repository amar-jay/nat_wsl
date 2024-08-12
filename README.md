Hi,
I recently ran into a problem when trying to access my server, which is running on my WSL Ubuntu distro, through my Wi-Fi router's IP address. The thing is, WSL doesn't automatically allow access through the router. Instead, it uses NAT (Network Address Translation) to route traffic through Windows, which is how I can access the internet from within WSL.

I'd like to set up a more comprehensive and straightforward NAT configuration. Ideally, I'd want to define the ports and settings for my WSL distro in a simple YAML file that automatically starts up when my system boots.

One way to simplify this process would be to run a script that calls netsh, which is the command-line tool Windows uses for managing network configurations. This script could handle all the NAT settings at runtime.

However, we're aiming to build a Network Address Translation system from scratch because we find it challenging and want to understand how it works in depth. By doing so, we'll gain hands-on experience with network fundamentals and have better control over our setup

Let's see if we can do so!!

---

### Docs

- [RFC 6866 - Port forwarding protocol](https://datatracker.ietf.org/doc/html/rfc6886)
