Pod::Spec.new do |spec|
  spec.name         = 'Gepv'
  spec.version      = '{{.Version}}'
  spec.license      = { :type => 'GNU Lesser General Public License, Version 3.0' }
  spec.homepage     = 'https://github.com/epvchain/go-epvchain'
  spec.authors      = { {{range .Contributors}}
		'{{.Name}}' => '{{.Email}}',{{end}}
	}
  spec.summary      = 'iOS EPVchain Client'
  spec.source       = { :git => 'https://github.com/epvchain/go-epvchain.git', :commit => '{{.Commit}}' }

	spec.platform = :ios
  spec.ios.deployment_target  = '9.0'
	spec.ios.vendored_frameworks = 'Frameworks/Gepv.framework'

	spec.prepare_command = <<-CMD
    curl https://gepvstore.blob.core.windows.net/builds/{{.Archive}}.tar.gz | tar -xvz
    mkdir Frameworks
    mv {{.Archive}}/Gepv.framework Frameworks
    rm -rf {{.Archive}}
  CMD
end
