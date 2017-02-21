# Config 1 - Single VM,Generic Cluster,  Ubuntu 14.04
$boxes = [
  {
    :cluster_size => 1,
    :ram => "2048",
    :os => "ubuntu",
    :os_version => "14.04",
    :ansible =>
      [{ :group => "node", :roles => [ "*" ]}],
    :forwarded_ports =>
      { 80 => 8480,
        443 => 8443,
        3306 => 3306,
        8080 => 8080,
        9001 => 9001,
        9002 => 9002,
        9003 => 9003,
        9004 => 9004
      },
  }
]
