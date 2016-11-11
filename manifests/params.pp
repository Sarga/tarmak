# == Class vault_client::params
#
# This class is meant to be called from vault_client.
# It sets variables according to platform.
#
class vault_client::params {
  $version = '0.6.2'
  $bin_dir = '/opt/bin'
  $dest_dir = '/opt'
  $config_dir = '/etc/vault'
  $download_dir = '/tmp'
  $download_url = 'https://releases.hashicorp.com/vault/#VERSION#/vault_#VERSION#_linux_amd64.zip'
}
